import { Step, StepIndicator, Stepper } from 'ui-components';

import { DFLink } from '@/components/DFLink';
import { InfoIcon } from '@/components/icons/common/Info';

export const GCPConnectorForm = () => {
  return (
    <Stepper>
      <Step
        indicator={
          <StepIndicator className="rounded-full">
            <div className="w-6 h-6 flex items-center justify-center">
              <span className="w-4 h-4">
                <InfoIcon />
              </span>
            </div>
          </StepIndicator>
        }
        title="Terraform"
      >
        <div className="text-p7 dark:text-text-text-and-icon">
          Connect to your Google Cloud Account via Terraform. Find out more information by{' '}
          <DFLink
            href={`https://docs.deepfence.io/threatmapper/docs/v2.0/cloudscanner/gcp`}
            target="_blank"
            rel="noreferrer"
            className="mt-2"
          >
            reading our documentation
          </DFLink>
          .
        </div>
      </Step>
    </Stepper>
  );
};
